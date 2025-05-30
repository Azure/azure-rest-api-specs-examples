
import com.azure.resourcemanager.support.models.Consent;
import com.azure.resourcemanager.support.models.ContactProfile;
import com.azure.resourcemanager.support.models.PreferredContactMethod;
import com.azure.resourcemanager.support.models.QuotaChangeRequest;
import com.azure.resourcemanager.support.models.QuotaTicketDetails;
import com.azure.resourcemanager.support.models.SeverityLevel;
import java.util.Arrays;

/**
 * Samples for SupportTickets Create.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/
     * CreateMachineLearningQuotaTicketForDedicatedCores.json
     */
    /**
     * Sample code: Create a ticket to request Quota increase for specific VM family cores for Machine Learning service.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void createATicketToRequestQuotaIncreaseForSpecificVMFamilyCoresForMachineLearningService(
        com.azure.resourcemanager.support.SupportManager manager) {
        manager.supportTickets().define("testticket").withDescription("my description").withProblemClassificationId(
            "/providers/Microsoft.Support/services/quota_service_guid/problemClassifications/machine_learning_service_problemClassification_guid")
            .withSeverity(SeverityLevel.MODERATE).withAdvancedDiagnosticConsent(Consent.YES)
            .withContactDetails(new ContactProfile().withFirstName("abc").withLastName("xyz")
                .withPreferredContactMethod(PreferredContactMethod.EMAIL).withPrimaryEmailAddress("abc@contoso.com")
                .withPreferredTimeZone("Pacific Standard Time").withCountry("usa")
                .withPreferredSupportLanguage("en-US"))
            .withTitle("my title").withServiceId("/providers/Microsoft.Support/services/quota_service_guid")
            .withSupportPlanId(
                "U291cmNlOlNDTSxDbGFyaWZ5SW5zdGFsbGF0aW9uU2l0ZUlkOjcsTGluZUl0ZW1JZDo5ODY1NzIyOSxDb250cmFjdElkOjk4NjU5MTk0LFN1YnNjcmlwdGlvbklkOjc2Y2I3N2ZhLThiMTctNGVhYi05NDkzLWI2NWRhY2U5OTgxMyw=")
            .withQuotaTicketDetails(new QuotaTicketDetails().withQuotaChangeRequestSubType("BatchAml")
                .withQuotaChangeRequestVersion("1.0")
                .withQuotaChangeRequests(Arrays.asList(new QuotaChangeRequest().withRegion("EastUS")
                    .withPayload("{\"VMFamily\":\"standardA0_A7Family\",\"NewLimit\":200,\"Type\":\"Dedicated\"}"))))
            .create();
    }
}
