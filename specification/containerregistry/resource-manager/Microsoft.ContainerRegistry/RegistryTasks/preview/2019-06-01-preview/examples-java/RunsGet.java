
/**
 * Samples for Runs Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/RegistryTasks/preview/2019-06-01-
     * preview/examples/RunsGet.json
     */
    /**
     * Sample code: Runs_Get.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void runsGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getRuns().getWithResponse("myResourceGroup", "myRegistry",
            "0accec26-d6de-4757-8e74-d080f38eaaab", com.azure.core.util.Context.NONE);
    }
}
