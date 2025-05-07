
import com.azure.resourcemanager.netapp.models.FilePathAvailabilityRequest;

/**
 * Samples for NetAppResource CheckFilePathAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/preview/2025-01-01-preview/examples/
     * CheckFilePathAvailability.json
     */
    /**
     * Sample code: CheckFilePathAvailability.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void checkFilePathAvailability(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.netAppResources().checkFilePathAvailabilityWithResponse("eastus",
            new FilePathAvailabilityRequest().withName("my-exact-filepth").withSubnetId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
            com.azure.core.util.Context.NONE);
    }
}
