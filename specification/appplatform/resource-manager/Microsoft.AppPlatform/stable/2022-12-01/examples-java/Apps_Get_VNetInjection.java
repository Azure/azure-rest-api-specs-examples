import com.azure.core.util.Context;

/** Samples for Apps Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/Apps_Get_VNetInjection.json
     */
    /**
     * Sample code: Apps_Get_VNetInjection.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void appsGetVNetInjection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getApps()
            .getWithResponse("myResourceGroup", "myservice", "myapp", null, Context.NONE);
    }
}
