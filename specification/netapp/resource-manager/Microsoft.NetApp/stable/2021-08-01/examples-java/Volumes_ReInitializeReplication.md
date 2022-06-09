```java
import com.azure.core.util.Context;

/** Samples for Volumes ReInitializeReplication. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2021-08-01/examples/Volumes_ReInitializeReplication.json
     */
    /**
     * Sample code: Volumes_ReInitializeReplication.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumesReInitializeReplication(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.volumes().reInitializeReplication("myRG", "account1", "pool1", "volume1", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-netapp_1.0.0-beta.7/sdk/netapp/azure-resourcemanager-netapp/README.md) on how to add the SDK to your project and authenticate.
