
import com.azure.resourcemanager.azurestackhci.models.DevicePoolProperties;

/**
 * Samples for DevicePools CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/DevicePools_CreateOrUpdate.json
     */
    /**
     * Sample code: DevicePools_CreateOrUpdate.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void devicePoolsCreateOrUpdate(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.devicePools().define("devicePool-1").withRegion("eastus").withExistingResourceGroup("ArcInstance-rg")
            .withProperties(new DevicePoolProperties()).create();
    }
}
