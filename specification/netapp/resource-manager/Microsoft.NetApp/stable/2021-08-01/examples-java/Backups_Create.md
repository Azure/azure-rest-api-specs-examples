Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-netapp_1.0.0-beta.7/sdk/netapp/azure-resourcemanager-netapp/README.md) on how to add the SDK to your project and authenticate.

```java
/** Samples for Backups Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2021-08-01/examples/Backups_Create.json
     */
    /**
     * Sample code: Backups_Create.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void backupsCreate(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager
            .backups()
            .define("backup1")
            .withRegion("eastus")
            .withExistingVolume("myRG", "account1", "pool1", "volume1")
            .withLabel("myLabel")
            .create();
    }
}
```
