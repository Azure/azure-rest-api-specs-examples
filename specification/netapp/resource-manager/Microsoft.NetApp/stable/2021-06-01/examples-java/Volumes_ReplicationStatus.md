Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-netapp_1.0.0-beta.6/sdk/netapp/azure-resourcemanager-netapp/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Volumes ReplicationStatus. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2021-06-01/examples/Volumes_ReplicationStatus.json
     */
    /**
     * Sample code: Volumes_ReplicationStatus.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumesReplicationStatus(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.volumes().replicationStatusWithResponse("myRG", "account1", "pool1", "volume1", Context.NONE);
    }
}
```
