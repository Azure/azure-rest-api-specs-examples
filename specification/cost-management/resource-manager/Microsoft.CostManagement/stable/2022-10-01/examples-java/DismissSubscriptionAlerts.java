import com.azure.resourcemanager.costmanagement.models.AlertStatus;
import com.azure.resourcemanager.costmanagement.models.DismissAlertPayload;

/** Samples for Alerts Dismiss. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/DismissSubscriptionAlerts.json
     */
    /**
     * Sample code: PatchSubscriptionAlerts.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void patchSubscriptionAlerts(com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .alerts()
            .dismissWithResponse(
                "subscriptions/00000000-0000-0000-0000-000000000000",
                "22222222-2222-2222-2222-222222222222",
                new DismissAlertPayload().withStatus(AlertStatus.DISMISSED),
                com.azure.core.util.Context.NONE);
    }
}
