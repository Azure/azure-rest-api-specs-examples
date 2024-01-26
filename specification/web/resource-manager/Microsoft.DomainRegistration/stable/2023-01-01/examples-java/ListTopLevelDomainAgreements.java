
import com.azure.resourcemanager.appservice.models.TopLevelDomainAgreementOption;

/**
 * Samples for TopLevelDomains ListAgreements.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.DomainRegistration/stable/2023-01-01/examples/
     * ListTopLevelDomainAgreements.json
     */
    /**
     * Sample code: List Top Level Domain Agreements.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listTopLevelDomainAgreements(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getTopLevelDomains().listAgreements("in",
            new TopLevelDomainAgreementOption().withIncludePrivacy(true).withForTransfer(false),
            com.azure.core.util.Context.NONE);
    }
}
