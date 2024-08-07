
import com.azure.resourcemanager.appcontainers.models.AccessMode;
import com.azure.resourcemanager.appcontainers.models.AzureFileProperties;
import com.azure.resourcemanager.appcontainers.models.ConnectedEnvironmentStorageProperties;

/**
 * Samples for ConnectedEnvironmentsStorages CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2024-03-01/examples/
     * ConnectedEnvironmentsStorages_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or update environments storage.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        createOrUpdateEnvironmentsStorage(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.connectedEnvironmentsStorages().define("jlaw-demo1")
            .withExistingConnectedEnvironment("examplerg", "env")
            .withProperties(new ConnectedEnvironmentStorageProperties().withAzureFile(
                new AzureFileProperties().withAccountName("account1").withAccountKey("fakeTokenPlaceholder")
                    .withAccessMode(AccessMode.READ_ONLY).withShareName("share1")))
            .create();
    }
}
