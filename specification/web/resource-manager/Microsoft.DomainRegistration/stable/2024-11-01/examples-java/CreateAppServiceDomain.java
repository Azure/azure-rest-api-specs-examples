
import com.azure.resourcemanager.appservice.fluent.models.DomainInner;
import com.azure.resourcemanager.appservice.models.Address;
import com.azure.resourcemanager.appservice.models.Contact;
import com.azure.resourcemanager.appservice.models.DnsType;
import com.azure.resourcemanager.appservice.models.DomainPurchaseConsent;
import java.time.OffsetDateTime;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Domains CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.DomainRegistration/stable/2024-11-01/examples/CreateAppServiceDomain
     * .json
     */
    /**
     * Sample code: Create App Service Domain.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAppServiceDomain(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getDomains().createOrUpdate("testrg123", "example.com",
            new DomainInner().withLocation("global").withTags(mapOf())
                .withContactAdmin(new Contact()
                    .withAddressMailing(new Address().withAddress1("3400 State St").withCity("Chicago")
                        .withCountry("United States").withPostalCode("fakeTokenPlaceholder").withState("IL"))
                    .withEmail("admin@email.com").withFax("1-245-534-2242").withJobTitle("Admin").withNameFirst("John")
                    .withNameLast("Doe").withNameMiddle("").withOrganization("Microsoft Inc.")
                    .withPhone("1-245-534-2242"))
                .withContactBilling(new Contact()
                    .withAddressMailing(new Address().withAddress1("3400 State St").withCity("Chicago")
                        .withCountry("United States").withPostalCode("fakeTokenPlaceholder").withState("IL"))
                    .withEmail("billing@email.com").withFax("1-245-534-2242").withJobTitle("Billing")
                    .withNameFirst("John").withNameLast("Doe").withNameMiddle("").withOrganization("Microsoft Inc.")
                    .withPhone("1-245-534-2242"))
                .withContactRegistrant(new Contact()
                    .withAddressMailing(new Address().withAddress1("3400 State St").withCity("Chicago")
                        .withCountry("United States").withPostalCode("fakeTokenPlaceholder").withState("IL"))
                    .withEmail("registrant@email.com").withFax("1-245-534-2242").withJobTitle("Registrant")
                    .withNameFirst("John").withNameLast("Doe").withNameMiddle("").withOrganization("Microsoft Inc.")
                    .withPhone("1-245-534-2242"))
                .withContactTech(new Contact()
                    .withAddressMailing(new Address().withAddress1("3400 State St").withCity("Chicago")
                        .withCountry("United States").withPostalCode("fakeTokenPlaceholder").withState("IL"))
                    .withEmail("tech@email.com").withFax("1-245-534-2242").withJobTitle("Tech").withNameFirst("John")
                    .withNameLast("Doe").withNameMiddle("").withOrganization("Microsoft Inc.")
                    .withPhone("1-245-534-2242"))
                .withPrivacy(false).withAutoRenew(true)
                .withConsent(new DomainPurchaseConsent().withAgreementKeys(Arrays.asList("agreementKey1"))
                    .withAgreedBy("192.0.2.1").withAgreedAt(OffsetDateTime.parse("2021-09-10T19:30:53Z")))
                .withDnsType(DnsType.DEFAULT_DOMAIN_REGISTRAR_DNS).withAuthCode("fakeTokenPlaceholder"),
            com.azure.core.util.Context.NONE);
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
