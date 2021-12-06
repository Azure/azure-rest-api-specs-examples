Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-netapp_1.0.0-beta.6/sdk/netapp/azure-resourcemanager-netapp/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.netapp.models.FilePathAvailabilityRequest;

/** Samples for NetAppResource CheckFilePathAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2021-06-01/examples/CheckFilePathAvailability.json
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
                Context.NONE);
    }
}
```
