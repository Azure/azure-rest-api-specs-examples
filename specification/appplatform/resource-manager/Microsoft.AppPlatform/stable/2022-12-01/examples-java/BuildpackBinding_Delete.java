import com.azure.core.util.Context;

/** Samples for BuildpackBinding Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/BuildpackBinding_Delete.json
     */
    /**
     * Sample code: BuildpackBinding_Delete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void buildpackBindingDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getBuildpackBindings()
            .delete("myResourceGroup", "myservice", "default", "default", "myBuildpackBinding", Context.NONE);
    }
}
