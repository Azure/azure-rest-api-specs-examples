
/**
 * Samples for ApplicationAccelerators Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/
     * ApplicationAccelerators_Delete.json
     */
    /**
     * Sample code: ApplicationAccelerators_Delete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void applicationAcceleratorsDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getApplicationAccelerators().delete("myResourceGroup",
            "myservice", "default", com.azure.core.util.Context.NONE);
    }
}
