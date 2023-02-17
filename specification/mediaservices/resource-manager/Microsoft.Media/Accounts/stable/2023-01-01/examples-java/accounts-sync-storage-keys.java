import com.azure.resourcemanager.mediaservices.models.SyncStorageKeysInput;

/** Samples for Mediaservices SyncStorageKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Accounts/stable/2023-01-01/examples/accounts-sync-storage-keys.json
     */
    /**
     * Sample code: Synchronizes Storage Account Keys.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void synchronizesStorageAccountKeys(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .mediaservices()
            .syncStorageKeysWithResponse(
                "contosorg",
                "contososports",
                new SyncStorageKeysInput().withId("contososportsstore"),
                com.azure.core.util.Context.NONE);
    }
}
