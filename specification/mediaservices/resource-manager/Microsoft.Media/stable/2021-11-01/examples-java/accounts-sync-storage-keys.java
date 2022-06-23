import com.azure.core.util.Context;
import com.azure.resourcemanager.mediaservices.models.SyncStorageKeysInput;

/** Samples for Mediaservices SyncStorageKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/accounts-sync-storage-keys.json
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
                "contoso", "contososports", new SyncStorageKeysInput().withId("contososportsstore"), Context.NONE);
    }
}
