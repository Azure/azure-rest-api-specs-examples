/** Samples for IngestionSettings ListTokens. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2021-01-15-preview/examples/IngestionSettings/ListTokensIngestionSetting_example.json
     */
    /**
     * Sample code: List ingestion setting tokens.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void listIngestionSettingTokens(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.ingestionSettings().listTokensWithResponse("default", com.azure.core.util.Context.NONE);
    }
}
