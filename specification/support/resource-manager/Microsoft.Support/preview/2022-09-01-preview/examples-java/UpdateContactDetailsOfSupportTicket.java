import com.azure.resourcemanager.support.models.PreferredContactMethod;
import com.azure.resourcemanager.support.models.UpdateContactProfile;
import com.azure.resourcemanager.support.models.UpdateSupportTicket;
import java.util.Arrays;

/** Samples for SupportTicketsNoSubscription Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2022-09-01-preview/examples/UpdateContactDetailsOfSupportTicket.json
     */
    /**
     * Sample code: Update contact details of a support ticket.
     *
     * @param manager Entry point to SupportManager.
     */
    public static void updateContactDetailsOfASupportTicket(com.azure.resourcemanager.support.SupportManager manager) {
        manager
            .supportTicketsNoSubscriptions()
            .updateWithResponse(
                "testticket",
                new UpdateSupportTicket()
                    .withContactDetails(
                        new UpdateContactProfile()
                            .withFirstName("first name")
                            .withLastName("last name")
                            .withPreferredContactMethod(PreferredContactMethod.EMAIL)
                            .withPrimaryEmailAddress("test.name@contoso.com")
                            .withAdditionalEmailAddresses(Arrays.asList("tname@contoso.com", "teamtest@contoso.com"))
                            .withPhoneNumber("123-456-7890")
                            .withPreferredTimeZone("Pacific Standard Time")
                            .withCountry("USA")
                            .withPreferredSupportLanguage("en-US")),
                com.azure.core.util.Context.NONE);
    }
}
