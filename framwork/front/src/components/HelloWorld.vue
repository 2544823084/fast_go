<script setup>
import { ref } from "vue";
import axios from "axios";

const uploading = ref(false);
const message = ref("");

const handleFileChange = async (event) => {
  const file = event.target.files?.[0];
  if (!file) return;

  const formData = new FormData();
  formData.append("file", file);

  uploading.value = true;
  message.value = "";
  try {
    // 不要手动设置 Content-Type，让浏览器/axios 自动带上 boundary
    const { data } = await axios.post(
      "http://localhost:3000/file/upload",
      formData,
    );
    message.value = `上传成功: ${data.filename} (${data.size} bytes)`;
  } catch (err) {
    message.value = `上传失败: ${err.response?.data?.error || err.message}`;
  } finally {
    uploading.value = false;
    event.target.value = "";
  }
};
</script>

<template>
  <div>
    <input type="file" :disabled="uploading" @change="handleFileChange" />
    <p v-if="uploading">上传中...</p>
    <p v-else-if="message">{{ message }}</p>
  </div>
</template>
