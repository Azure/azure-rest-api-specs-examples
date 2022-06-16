import com.azure.core.util.Context;

/** Samples for ManagedClusters GetCommandResult. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2022-04-01/examples/RunCommandResultFailed.json
     */
    /**
     * Sample code: commandFailedResult.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void commandFailedResult(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .kubernetesClusters()
            .manager()
            .serviceClient()
            .getManagedClusters()
            .getCommandResultWithResponse("rg1", "clustername1", "def7b3ea71bd4f7e9d226ddbc0f00ad9", Context.NONE);
    }
}
