
import com.azure.resourcemanager.support.models.ServiceClassificationRequest;

/**
 * Samples for ServiceClassificationsNoSubscription ClassifyServices.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/support/resource-manager/Microsoft.Support/preview/2023-06-01-preview/examples/ClassifyServices.
     * json
     */
    /**
     * Sample code: Classify list of Azure services.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void classifyListOfAzureServices(com.azure.resourcemanager.support.SupportManager manager) {
        manager.serviceClassificationsNoSubscriptions().classifyServicesWithResponse(
            new ServiceClassificationRequest().withIssueSummary("Can not connect to Windows VM").withResourceId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgname/providers/Microsoft.Compute/virtualMachines/vmname"),
            com.azure.core.util.Context.NONE);
    }
}
