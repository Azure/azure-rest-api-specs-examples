import com.azure.core.util.Context;

/** Samples for Runs Cancel. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/RunsCancel.json
     */
    /**
     * Sample code: Runs_Cancel.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void runsCancel(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getRuns()
            .cancel("myResourceGroup", "myRegistry", "0accec26-d6de-4757-8e74-d080f38eaaab", Context.NONE);
    }
}
