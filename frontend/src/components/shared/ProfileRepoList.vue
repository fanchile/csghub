<template>
  <div class="sm:w-[100%] sm:mt-[36px]">
    <!-- collections -->
    <div>
      <h3 class="text-xl text-gray-700 flex items-center gap-[8px]">
        <SvgIcon name="collections" width="18" height="18" />
        <span>{{ $t("collections.collection") }}</span>
      </h3>
      <div v-if="hasCollections" class="mb-4 mt-[16px]">
        <CollectionCards 
          v-if="collections.data.length > 0" 
          :collections="collections.data"
        />
      </div>
      <div v-else class="flex flex-wrap gap-4 mb-4 mt-[16px]">
        {{ $t("all.noData") }}
      </div>
      <view-more v-if="collections.more" target="collections" @view-more-targets="viewMoreTargets"></view-more>
      <el-skeleton class="pr-6" v-if="collectionsLoading" :rows="2" animated />
    </div>

    <!-- models -->
    <div>
      <h3 class="text-xl text-gray-700 flex items-center gap-[8px]">
        <SvgIcon name="models" width="18" height="18" />
        <span>{{ $t("organization.model") }}</span>
      </h3>
      <div v-if="hasModels" class="grid grid-cols-2 xl:grid-cols-1 gap-4 mb-4 mt-[16px]">
        <repo-item v-for="model in models.data" :repo="model" repo-type="model"></repo-item>
      </div>
      <div v-else class="flex flex-wrap gap-4 mb-4 mt-[16px]">
        {{ $t("all.noData") }}
      </div>
      <view-more v-if="models.more" target="models" @view-more-targets="viewMoreTargets"></view-more>
      <el-skeleton class="pr-6" v-if="modelsLoading" :rows="2" animated />
    </div>

    <!-- datasets -->
    <div class="mt-[32px]">
      <h3 class="text-xl text-gray-700 flex items-center gap-[8px]">
        <SvgIcon name="datasets" width="18" height="18" />
        <span>{{ $t("organization.dataset") }}</span>
      </h3>
      <div v-if="hasDatasets" class="grid grid-cols-2 xl:grid-cols-1 gap-4 mb-4 mt-[16px]">
        <repo-item v-for="dataset in datasets.data" :repo="dataset" repo-type="dataset"></repo-item>
      </div>
      <div v-else class="flex flex-wrap gap-4 mb-4 mt-[16px]">
        {{ $t("all.noData") }}
      </div>
      <view-more
        v-if="datasets.more"
        target="datasets"
        @view-more-targets="viewMoreTargets"
      ></view-more>
      <el-skeleton class="pr-6" v-if="datasetsLoading" :rows="2" animated />
    </div>

    <!-- code repo -->
    <div class="mt-[32px]">
      <h3 class="text-xl text-gray-700 flex items-center gap-[8px]">
        <SvgIcon name="codes" width="18" height="18" />
        <span>{{ $t("organization.code") }}</span>
      </h3>
      <div v-if="hasCodes" class="grid grid-cols-2 xl:grid-cols-1 gap-4 mb-4 mt-[16px]">
        <repo-item v-for="code in codes.data" :repo="code" repo-type="code"></repo-item>
      </div>
      <div v-else class="flex flex-wrap gap-4 mb-4 mt-[16px]">
        {{ $t("all.noData") }}
      </div>
      <view-more v-if="codes.more" target="codes" @view-more-targets="viewMoreTargets"></view-more>
      <el-skeleton class="pr-6" v-if="codeLoading" :rows="2" animated />
    </div>

    <!-- spaces -->
    <div class="mt-[32px]">
      <h3 class="text-xl text-gray-700 flex items-center gap-[8px]">
        <SvgIcon name="spaces" width="18" height="18" />
        <span>{{ $t("organization.space") }}</span>
      </h3>
      <div v-if="hasSpaces" class="grid grid-cols-2 xl:grid-cols-1 gap-4 mb-4 mt-[16px]">
        <application-space-item v-for="repo in spaces.data" :repo="repo" repo-type="space" />
      </div>
      <div v-else class="flex flex-wrap gap-4 mb-4 mt-[16px]">
        {{ $t("all.noData") }}
      </div>
      <view-more v-if="spaces.more" target="spaces" @view-more-targets="viewMoreTargets"></view-more>
      <el-skeleton class="pr-6" v-if="spacesLoading" :rows="2" animated />
    </div>
  </div>
</template>

<script setup>
  import { computed, ref, onMounted } from "vue"
  import RepoItem from "./RepoItem.vue"
  import CollectionCards from "../collections/CollectionCards.vue"
  import ApplicationSpaceItem from "../application_spaces/ApplicationSpaceItem.vue"
  import ViewMore from "./ViewMore.vue"
  import { useI18n } from "vue-i18n"
  import useFetchApi from "../../packs/useFetchApi"
  import { ElMessage } from "element-plus"
  import { useCookies } from 'vue3-cookies'
  const { cookies } = useCookies()
  const props = defineProps({
    name: String,
    initiator: String
  })

  const { t } = useI18n()
  const current_user = cookies.get('current_user')
  const isCurrentUser = computed(() => {
    return props.name === current_user
  })

  const collections = ref([])
  const models = ref([])
  const datasets = ref([])
  const codes = ref([])
  const spaces = ref([])
  const collectionsLoading = ref(false)
  const modelsLoading = ref(false)
  const datasetsLoading = ref(false)
  const codeLoading = ref(false)
  const spacesLoading = ref(false)

  const hasCollections = computed(() => collections.value?.total > 0)
  const hasModels = computed(() => models.value?.total > 0)
  const hasDatasets = computed(() => datasets.value?.total > 0)
  const hasCodes = computed(() => codes.value?.total > 0)
  const hasSpaces = computed(() => spaces.value?.total > 0)

  const prefixPath =
    document.location.pathname.split("/")[1] === "organizations" ? "organization" : "user"

  const getProfileRepoData = async () =>{
    const defaultTotal = 6
    const collectionsUrl = reposUrl("collections")
    const modelsUrl = reposUrl("models")
    const datasetsUrl = reposUrl("datasets")
    const spacesUrl = reposUrl("spaces")
    const codesUrl = reposUrl("codes")
    const promises = [
        fetchData(collectionsUrl, collections, defaultTotal),
        fetchData(modelsUrl, models, defaultTotal),
        fetchData(datasetsUrl, datasets, defaultTotal),
        fetchData(spacesUrl, spaces, defaultTotal),
        fetchData(codesUrl, codes, defaultTotal)
    ];
    await Promise.all(promises);
  }

  const viewMoreTargets = (target) => {
    if (target === "models") {
      modelsLoading.value = true
      fetchMoreModels()
    } else if (target === "datasets") {
      datasetsLoading.value = true
      fetchMoreDatasets()
    } else if (target === "spaces") {
      spacesLoading.value = true
      fetchMoreSpaces()
    } else if (target === "codes") {
      codeLoading.value = true
      fetchMoreCodes()
    } else if (target === "collections") {
      collectionsLoading.value = true
      fetchMoreCollections()
    }
  }

  const reposUrl = (type) => {
    switch (props.initiator) {
      case "likes":
        return `/${prefixPath}/${props.name}/likes/${type}`
      case "profile":
      default:
        return `/${prefixPath}/${props.name}/${type}`
    }
  }

  const fetchMoreCollections = async () => {
    const url = reposUrl("collections")
    await fetchData(url, collections, collections.value.total)
  }
  const fetchMoreModels = async () => {
    const url = reposUrl("models")
    await fetchData(url, models, models.value.total)
  }

  const fetchMoreDatasets = async () => {
    const url = reposUrl("datasets")
    await fetchData(url, datasets, datasets.value.total)
  }

  const fetchMoreSpaces = async () => {
    const url = reposUrl("spaces")
    await fetchData(url, spaces, spaces.value.total)
  }

  const fetchMoreCodes = async () => {
    const url = reposUrl("codes")
    await fetchData(url, codes, codes.value.total)
  }

  const fetchData = async (url, targetRef, total, type) => {
    const params = new URLSearchParams()
    params.append("per", total)
    params.append("page", 1)
    
    try {
      const { data, error } = await useFetchApi(`${url}?${params}`).json()

      if (error.value) {
        ElMessage({
          message: error.value.msg,
          type: "warning"
        })
      } else {
        targetRef.value = data.value
        if(targetRef.value?.total > 6 && total === 6){
          targetRef.value.more = true
        }else if(total > 6){
          targetRef.value.more = false
        }
      }
    } catch (error) {
      console.log(error)
    } finally {
      if (targetRef === models) {
        modelsLoading.value = false
      } else if (targetRef === datasets) {
        datasetsLoading.value = false
      } else if (targetRef === codes) {
        codeLoading.value = false
      } else if (targetRef === spaces) {
        spacesLoading.value = false
      }
    }
  }

  onMounted(() => {
    getProfileRepoData()
  })
</script>
