
/**
 * Samples for UpdateRuns Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/
     * GetUpdateRuns.json
     */
    /**
     * Sample code: Get Update runs under cluster resource.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        getUpdateRunsUnderClusterResource(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.updateRuns().getWithResponse("testrg", "testcluster", "Microsoft4.2203.2.32",
            "23b779ba-0d52-4a80-8571-45ca74664ec3", com.azure.core.util.Context.NONE);
    }
}
