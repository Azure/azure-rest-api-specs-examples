/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/datadog/resource-manager/Microsoft.Datadog/stable/2023-01-01/examples/Operations_List.json
     */
    /**
     * Sample code: Operations_List.
     *
     * @param manager Entry point to MicrosoftDatadogManager.
     */
    public static void operationsList(com.azure.resourcemanager.datadog.MicrosoftDatadogManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
