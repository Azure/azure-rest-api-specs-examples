
import com.azure.resourcemanager.communication.models.DomainResource;
import com.azure.resourcemanager.communication.models.UserEngagementTracking;

/**
 * Samples for Domains Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/communication/resource-manager/Microsoft.Communication/stable/2025-09-01/examples/domains/update.
     * json
     */
    /**
     * Sample code: Update Domains resource.
     * 
     * @param manager Entry point to CommunicationManager.
     */
    public static void updateDomainsResource(com.azure.resourcemanager.communication.CommunicationManager manager) {
        DomainResource resource = manager.domains().getWithResponse("MyResourceGroup", "MyEmailServiceResource",
            "mydomain.com", com.azure.core.util.Context.NONE).getValue();
        resource.update().withUserEngagementTracking(UserEngagementTracking.ENABLED).apply();
    }
}
