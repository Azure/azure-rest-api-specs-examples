
/**
 * Samples for ManagedClusters GetCommandResult.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/
     * RunCommandResultSucceed.json
     */
    /**
     * Sample code: commandSucceedResult.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void commandSucceedResult(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getManagedClusters().getCommandResultWithResponse("rg1",
            "clustername1", "def7b3ea71bd4f7e9d226ddbc0f00ad9", com.azure.core.util.Context.NONE);
    }
}
