
/**
 * Samples for Assets Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01-preview/Delete_Asset.json
     */
    /**
     * Sample code: Delete_Asset.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void deleteAsset(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.assets().delete("myResourceGroup", "my-asset", com.azure.core.util.Context.NONE);
    }
}
