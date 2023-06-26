import com.azure.core.util.Context;

/** Samples for Apps List. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/Apps_List_VNetInjection.json
     */
    /**
     * Sample code: Apps_List_VNetInjection.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void appsListVNetInjection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getApps().list("myResourceGroup", "myservice", Context.NONE);
    }
}
