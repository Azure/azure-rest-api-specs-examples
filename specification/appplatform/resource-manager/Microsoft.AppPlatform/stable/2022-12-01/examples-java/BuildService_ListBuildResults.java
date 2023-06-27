import com.azure.core.util.Context;

/** Samples for BuildService ListBuildResults. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/BuildService_ListBuildResults.json
     */
    /**
     * Sample code: BuildService_ListBuildResults.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void buildServiceListBuildResults(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getBuildServices()
            .listBuildResults("myResourceGroup", "myservice", "default", "mybuild", Context.NONE);
    }
}
