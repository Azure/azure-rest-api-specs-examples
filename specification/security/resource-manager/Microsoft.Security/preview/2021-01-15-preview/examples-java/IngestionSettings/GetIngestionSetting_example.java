import com.azure.core.util.Context;

/** Samples for IngestionSettings Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2021-01-15-preview/examples/IngestionSettings/GetIngestionSetting_example.json
     */
    /**
     * Sample code: Get a ingestion setting on subscription.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void getAIngestionSettingOnSubscription(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.ingestionSettings().getWithResponse("default", Context.NONE);
    }
}
