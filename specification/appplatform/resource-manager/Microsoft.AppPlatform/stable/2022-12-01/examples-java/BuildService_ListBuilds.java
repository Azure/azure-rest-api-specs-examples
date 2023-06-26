import com.azure.core.util.Context;

/** Samples for BuildService ListBuilds. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/BuildService_ListBuilds.json
     */
    /**
     * Sample code: BuildService_ListBuilds.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void buildServiceListBuilds(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getBuildServices()
            .listBuilds("myResourceGroup", "myservice", "default", Context.NONE);
    }
}
