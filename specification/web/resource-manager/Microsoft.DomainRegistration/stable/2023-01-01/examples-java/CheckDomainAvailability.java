
import com.azure.resourcemanager.appservice.fluent.models.NameIdentifierInner;

/**
 * Samples for Domains CheckAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.DomainRegistration/stable/2023-01-01/examples/
     * CheckDomainAvailability.json
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
