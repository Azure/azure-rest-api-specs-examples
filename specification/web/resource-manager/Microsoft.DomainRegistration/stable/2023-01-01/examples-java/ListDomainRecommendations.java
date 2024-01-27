
import com.azure.resourcemanager.appservice.models.DomainRecommendationSearchParameters;

/**
 * Samples for Domains ListRecommendations.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.DomainRegistration/stable/2023-01-01/examples/
     * ListDomainRecommendations.json
     */
    /**
     * Sample code: List domain recommendations.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listDomainRecommendations(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getDomains()
            .listRecommendations(new DomainRecommendationSearchParameters().withKeywords("fakeTokenPlaceholder")
                .withMaxDomainRecommendations(10), com.azure.core.util.Context.NONE);
    }
}
