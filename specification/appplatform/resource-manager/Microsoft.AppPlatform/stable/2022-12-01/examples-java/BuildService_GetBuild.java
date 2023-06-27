import com.azure.core.util.Context;

/** Samples for BuildService GetBuild. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/BuildService_GetBuild.json
     */
    /**
     * Sample code: BuildService_GetBuild.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void buildServiceGetBuild(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getBuildServices()
            .getBuildWithResponse("myResourceGroup", "myservice", "default", "mybuild", Context.NONE);
    }
}
