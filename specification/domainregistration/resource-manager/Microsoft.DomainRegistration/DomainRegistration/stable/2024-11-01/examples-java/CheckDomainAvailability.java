
import com.azure.resourcemanager.appservice.fluent.models.NameIdentifierInner;

/**
 * Samples for Domains CheckAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/domainregistration/resource-manager/Microsoft.DomainRegistration/DomainRegistration/stable/2024-11-
     * 01/examples/CheckDomainAvailability.json
     */
    /**
     * Sample code: Check domain availability.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void checkDomainAvailability(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getDomains().checkAvailabilityWithResponse(
            new NameIdentifierInner().withName("abcd.com"), com.azure.core.util.Context.NONE);
    }
}
