/** Samples for IngestionSettings Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2021-01-15-preview/examples/IngestionSettings/CreateIngestionSetting_example.json
     */
    /**
     * Sample code: Create an ingestion setting for subscription.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void createAnIngestionSettingForSubscription(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.ingestionSettings().define("default").create();
    }
}
