
import com.azure.resourcemanager.neonpostgres.models.CompanyDetails;
import com.azure.resourcemanager.neonpostgres.models.OrganizationProperties;
import com.azure.resourcemanager.neonpostgres.models.OrganizationResource;
import com.azure.resourcemanager.neonpostgres.models.PartnerOrganizationProperties;
import com.azure.resourcemanager.neonpostgres.models.SingleSignOnProperties;
import com.azure.resourcemanager.neonpostgres.models.SingleSignOnStates;
import com.azure.resourcemanager.neonpostgres.models.UserDetails;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Organizations Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-08-01-preview/Organizations_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: Organizations_Update.
     * 
     * @param manager Entry point to NeonPostgresManager.
     */
    public static void organizationsUpdate(com.azure.resourcemanager.neonpostgres.NeonPostgresManager manager) {
        OrganizationResource resource = manager.organizations()
            .getByResourceGroupWithResponse("rgneon", "eRY-J_:", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("key8990", "fakeTokenPlaceholder"))
            .withProperties(new OrganizationProperties()
                .withUserDetails(new UserDetails().withFirstName("buwwe").withLastName("escynjpynkoox")
                    .withEmailAddress("3i_%@w8-y.H-p.tvj.dG").withUpn("fwedjamgwwrotcjaucuzdwycfjdqn")
                    .withPhoneNumber("dlrqoowumy"))
                .withCompanyDetails(new CompanyDetails().withCompanyName("uxn").withCountry("lpajqzptqchuko")
                    .withOfficeAddress("chpkrlpmfslmawgunjxdllzcrctykq").withBusinessPhone("hbeb")
                    .withDomain("krjldeakhwiepvs").withNumberOfEmployees(23L))
                .withPartnerOrganizationProperties(new PartnerOrganizationProperties()
                    .withOrganizationId("njyoqflcmfwzfsqe").withOrganizationName("J:.._3P")
                    .withSingleSignOnProperties(new SingleSignOnProperties()
                        .withSingleSignOnState(SingleSignOnStates.INITIAL).withEnterpriseAppId("fpibacregjfncfdsojs")
                        .withSingleSignOnUrl("tmojh").withAadDomains(Arrays.asList("kndszgrwzbvvlssvkej")))))
            .apply();
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
