/** Samples for IngestionSettings ListConnectionStrings. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2021-01-15-preview/examples/IngestionSettings/ListConnectionStrings_example.json
     */
    /**
     * Sample code: List connection strings for ingesting security data and logs.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void listConnectionStringsForIngestingSecurityDataAndLogs(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.ingestionSettings().listConnectionStringsWithResponse("default", com.azure.core.util.Context.NONE);
    }
}
