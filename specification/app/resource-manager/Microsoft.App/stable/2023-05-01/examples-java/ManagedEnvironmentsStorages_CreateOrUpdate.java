import com.azure.resourcemanager.appcontainers.models.AccessMode;
import com.azure.resourcemanager.appcontainers.models.AzureFileProperties;
import com.azure.resourcemanager.appcontainers.models.ManagedEnvironmentStorageProperties;

/** Samples for ManagedEnvironmentsStorages CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2023-05-01/examples/ManagedEnvironmentsStorages_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or update environments storage.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void createOrUpdateEnvironmentsStorage(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager
            .managedEnvironmentsStorages()
            .define("jlaw-demo1")
            .withExistingManagedEnvironment("examplerg", "managedEnv")
            .withProperties(
                new ManagedEnvironmentStorageProperties()
                    .withAzureFile(
                        new AzureFileProperties()
                            .withAccountName("account1")
                            .withAccountKey("fakeTokenPlaceholder")
                            .withAccessMode(AccessMode.READ_ONLY)
                            .withShareName("share1")))
            .create();
    }
}
