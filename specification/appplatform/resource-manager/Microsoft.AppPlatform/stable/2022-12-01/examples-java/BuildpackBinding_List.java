import com.azure.core.util.Context;

/** Samples for BuildpackBinding List. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/BuildpackBinding_List.json
     */
    /**
     * Sample code: BuildpackBinding_Get.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void buildpackBindingGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getBuildpackBindings()
            .list("myResourceGroup", "myservice", "default", "default", Context.NONE);
    }
}
