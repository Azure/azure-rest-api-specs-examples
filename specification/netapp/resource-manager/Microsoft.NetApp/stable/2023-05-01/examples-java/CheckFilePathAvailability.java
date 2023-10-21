import com.azure.resourcemanager.netapp.models.FilePathAvailabilityRequest;

/** Samples for NetAppResource CheckFilePathAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2023-05-01/examples/CheckFilePathAvailability.json
     */
    /**
     * Sample code: CheckFilePathAvailability.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void checkFilePathAvailability(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager
            .netAppResources()
            .checkFilePathAvailabilityWithResponse(
                "eastus",
                new FilePathAvailabilityRequest()
                    .withName("my-exact-filepth")
                    .withSubnetId(
                        "/subscriptions/9760acf5-4638-11e7-9bdb-020073ca7778/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
                com.azure.core.util.Context.NONE);
    }
}
