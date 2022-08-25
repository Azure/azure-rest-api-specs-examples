import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2019-05-05-preview/examples/Operations_List.json
     */
    /**
     * Sample code: ListOperations.
     *
     * @param manager Entry point to AlertsManagementManager.
     */
    public static void listOperations(com.azure.resourcemanager.alertsmanagement.AlertsManagementManager manager) {
        manager.operations().list(Context.NONE);
    }
}
