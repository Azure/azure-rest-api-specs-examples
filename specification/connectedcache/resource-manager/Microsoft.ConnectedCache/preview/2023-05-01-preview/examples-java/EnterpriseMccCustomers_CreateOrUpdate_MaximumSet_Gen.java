
import com.azure.resourcemanager.connectedcache.models.AdditionalCustomerProperties;
import com.azure.resourcemanager.connectedcache.models.CustomerEntity;
import com.azure.resourcemanager.connectedcache.models.CustomerProperty;
import com.azure.resourcemanager.connectedcache.models.CustomerTransitState;
import java.time.OffsetDateTime;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for EnterpriseMccCustomers CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-05-01-preview/EnterpriseMccCustomers_CreateOrUpdate_MaximumSet_Gen.json
     */
    /**
     * Sample code: EnterpriseMccCustomers_CreateOrUpdate.
     * 
     * @param manager Entry point to ConnectedCacheManager.
     */
    public static void
        enterpriseMccCustomersCreateOrUpdate(com.azure.resourcemanager.connectedcache.ConnectedCacheManager manager) {
        manager.enterpriseMccCustomers().define("MccRPTest1").withRegion("westus")
            .withExistingResourceGroup("rgConnectedCache").withTags(mapOf("key3379", "fakeTokenPlaceholder"))
            .withProperties(new CustomerProperty()
                .withCustomer(new CustomerEntity().withFullyQualifiedResourceId("uqsbtgae")
                    .withCustomerName("mkpzynfqihnjfdbaqbqwyhd").withContactEmail("xquos").withContactPhone("vue")
                    .withContactName("wxyqjoyoscmvimgwhpitxky").withIsEntitled(true).withReleaseVersion(20)
                    .withClientTenantId("fproidkpgvpdnac").withIsEnterpriseManaged(true).withShouldMigrate(true)
                    .withResendSignupCode(true).withVerifySignupCode(true)
                    .withVerifySignupPhrase("tprjvttkgmrqlsyicnidhm"))
                .withAdditionalCustomerProperties(new AdditionalCustomerProperties().withCustomerEmail("zdjgibsidydyzm")
                    .withCustomerTransitAsn("habgklnxqzmozqpazoyejwiphezpi")
                    .withCustomerTransitState(CustomerTransitState.fromString("voblixkxfejbmhxilb"))
                    .withCustomerAsn("hgrelgnrtdkleisnepfolu").withCustomerEntitlementSkuId("b")
                    .withCustomerEntitlementSkuGuid("rvzmdpxyflgqetvpwupnfaxsweiiz")
                    .withCustomerEntitlementSkuName("waaqfijr")
                    .withCustomerEntitlementExpiration(OffsetDateTime.parse("2024-01-30T00:54:04.773Z"))
                    .withOptionalProperty1("qhmwxza").withOptionalProperty2("l").withOptionalProperty3("mblwwvbie")
                    .withOptionalProperty4("vzuek").withOptionalProperty5("fzjodscdfcdr")))
            .create();
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
