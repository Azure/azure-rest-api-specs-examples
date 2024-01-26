
/**
 * Samples for RuntimeVersions ListRuntimeVersions.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/
     * RuntimeVersions_ListRuntimeVersions.json
     */
    /**
     * Sample code: RuntimeVersions_ListRuntimeVersions.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void runtimeVersionsListRuntimeVersions(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getRuntimeVersions()
            .listRuntimeVersionsWithResponse(com.azure.core.util.Context.NONE);
    }
}
