
/**
 * Samples for Usages ListByLocation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/StorageAccountListLocationUsage.json
     */
    /**
     * Sample code: UsageList.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void usageList(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getUsages().listByLocation("eastus2(stage)", com.azure.core.util.Context.NONE);
    }
}
