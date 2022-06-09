```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.netapp.models.AuthorizeRequest;

/** Samples for Volumes AuthorizeReplication. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2021-06-01/examples/Volumes_AuthorizeReplication.json
     */
    /**
     * Sample code: Volumes_AuthorizeReplication.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumesAuthorizeReplication(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager
            .volumes()
            .authorizeReplication(
                "myRG",
                "account1",
                "pool1",
                "volume1",
                new AuthorizeRequest()
                    .withRemoteVolumeResourceId(
                        "/subscriptions/D633CC2E-722B-4AE1-B636-BBD9E4C60ED9/resourceGroups/myRemoteRG/providers/Microsoft.NetApp/netAppAccounts/remoteAccount1/capacityPools/remotePool1/volumes/remoteVolume1"),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-netapp_1.0.0-beta.6/sdk/netapp/azure-resourcemanager-netapp/README.md) on how to add the SDK to your project and authenticate.
