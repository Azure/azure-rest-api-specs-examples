```java
import com.azure.resourcemanager.recoveryservicesbackup.models.BackupManagementType;
import com.azure.resourcemanager.recoveryservicesbackup.models.ProtectionContainer;

/** Samples for ProtectionContainers Register. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-07-01/examples/AzureStorage/ProtectionContainers_Register.json
     */
    /**
     * Sample code: RegisterAzure Storage ProtectionContainers.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void registerAzureStorageProtectionContainers(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .protectionContainers()
            .define("VMAppContainer;Compute;testRG;testSQL")
            .withRegion((String) null)
            .withExistingBackupFabric("testvault", "test-rg", "Azure")
            .withProperties(
                new ProtectionContainer()
                    .withFriendlyName("testSQL")
                    .withBackupManagementType(BackupManagementType.AZURE_WORKLOAD))
            .create();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-recoveryservicesbackup_1.0.0-beta.2/sdk/recoveryservicesbackup/azure-resourcemanager-recoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.
