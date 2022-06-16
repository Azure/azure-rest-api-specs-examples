import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/stable/2020-05-01/examples/Operations_List.json
     */
    /**
     * Sample code: GetHealthHistoryByResource.
     *
     * @param manager Entry point to ResourceHealthManager.
     */
    public static void getHealthHistoryByResource(
        com.azure.resourcemanager.resourcehealth.ResourceHealthManager manager) {
        manager.operations().listWithResponse(Context.NONE);
    }
}
