<script setup lang="ts">
import { reactive, onMounted } from "vue";
import { useCounterService } from "./CounterService";
import type { Response } from "@/types/response";

type State = {
  count: number;
  loading: boolean;
  error: string;
};

const counterService = useCounterService();

const state = reactive<State>({
  count: 0,
  loading: false,
  error: "",
});

const setState = (newState: State) => {
  state.count = newState.count;
  state.loading = newState.loading;
  state.error = newState.error;
};

const makeRequest = (cb: () => Promise<Response<number>>): void => {
  setState({ ...state, loading: true });
  cb()
    .then((res) => {
      if (res.success) {
        setState({ ...state, count: res.data!, loading: false });
      } else {
        setState({ ...state, error: res.error!, loading: false });
      }
    })
    .catch((err) => {
      setState({ ...state, error: err.message, loading: false });
    });
};

const getCount = (): void => {
  makeRequest(() => counterService.get());
};

const setCount = (count: number): void => {
  makeRequest(() => counterService.set(count));
};

const incrementCount = (): void => {
  makeRequest(() => counterService.increment());
};

const decrementCount = (): void => {
  makeRequest(() => counterService.decrement());
};

const resetCount = (): void => {
  makeRequest(() => counterService.reset());
};

onMounted(() => {
  getCount();
});
</script>

<template>
  <div class="counter-view">
    <h1>Counter</h1>
    <template v-if="state.loading">
      <p>Loading...</p>
    </template>
    <template v-else-if="state.error">
      <p>{{ state.error }}</p>
    </template>
    <p>Count: {{ state.count }}</p>
    <button class="btn" @click="incrementCount" :disabled="state.loading">
      Increment
    </button>
    <button class="btn" @click="decrementCount" :disabled="state.loading">
      Decrement
    </button>
    <button class="btn" @click="resetCount" :disabled="state.loading">
      Reset
    </button>
    <button class="btn" @click="setCount(100)" :disabled="state.loading">
      Set to 100
    </button>
  </div>
</template>

<style lang="scss" scoped>
.counter-view {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100%;
  width: 100%;
  padding: 1rem;
  box-sizing: border-box;
  h1 {
    margin: 0;
    font-size: 2rem;
  }
  p {
    margin: 0;
    font-size: 1.5rem;
  }
  .btn {
    margin: 0.5rem;
    padding: 0.5rem 1rem;
    font-size: 1rem;
    border: none;
    border-radius: 0.25rem;
    background-color: var(--vt-c-text-light-1);
    color: var(--vt-c-white);
    cursor: pointer;
    &:disabled {
      opacity: 0.5;
      cursor: not-allowed;
    }
    &:active {
      opacity: 0.8;
    }
  }
}
</style>
