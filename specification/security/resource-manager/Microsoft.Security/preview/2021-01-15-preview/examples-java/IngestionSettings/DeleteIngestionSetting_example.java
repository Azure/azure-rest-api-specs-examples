/** Samples for IngestionSettings Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2021-01-15-preview/examples/IngestionSettings/DeleteIngestionSetting_example.json
     */
    /**
     * Sample code: Delete an ingestion setting for the subscription.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void deleteAnIngestionSettingForTheSubscription(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.ingestionSettings().deleteWithResponse("default", com.azure.core.util.Context.NONE);
    }
}
