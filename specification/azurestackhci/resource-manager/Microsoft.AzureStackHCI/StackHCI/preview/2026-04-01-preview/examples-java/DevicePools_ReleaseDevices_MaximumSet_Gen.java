
import com.azure.resourcemanager.azurestackhci.models.ReleaseDeviceRequest;
import java.util.Arrays;

/**
 * Samples for DevicePools ReleaseDevices.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/DevicePools_ReleaseDevices_MaximumSet_Gen.json
     */
    /**
     * Sample code: DevicePools_ReleaseDevices_MaximumSet.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        devicePoolsReleaseDevicesMaximumSet(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.devicePools().releaseDevices("ArcInstance-rg", "snbyzreoirqiz",
            new ReleaseDeviceRequest().withDevices(Arrays.asList(
                "/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.AzureStackHCI/edgeMachines/machine-1")),
            com.azure.core.util.Context.NONE);
    }
}
