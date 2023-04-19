import com.azure.resourcemanager.qumulo.models.ManagedServiceIdentity;
import com.azure.resourcemanager.qumulo.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.qumulo.models.MarketplaceDetails;
import com.azure.resourcemanager.qumulo.models.StorageSku;
import com.azure.resourcemanager.qumulo.models.UserAssignedIdentity;
import com.azure.resourcemanager.qumulo.models.UserDetails;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for FileSystems CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/liftrqumulo/resource-manager/Qumulo.Storage/preview/2022-10-12-preview/examples/FileSystems_CreateOrUpdate_MaximumSet_Gen.json
     */
    /**
     * Sample code: FileSystems_CreateOrUpdate_MaximumSet_Gen.
     *
     * @param manager Entry point to QumuloManager.
     */
    public static void fileSystemsCreateOrUpdateMaximumSetGen(com.azure.resourcemanager.qumulo.QumuloManager manager) {
        manager
            .fileSystems()
            .define("nauwwbfoqehgbhdsmkewoboyxeqg")
            .withRegion("przdlsmlzsszphnixq")
            .withExistingResourceGroup("rgQumulo")
            .withMarketplaceDetails(
                new MarketplaceDetails()
                    .withMarketplaceSubscriptionId("ujrcqvxfnhxxheoth")
                    .withPlanId("x")
                    .withOfferId("eiyhbmpwgezcmzrrfoiskuxlcvwojf")
                    .withPublisherId("wfmokfdjbwpjhz"))
            .withStorageSku(StorageSku.STANDARD)
            .withUserDetails(new UserDetails().withEmail("viptslwulnpaupfljvnjeq"))
            .withDelegatedSubnetId("neqctctqdmjezfgt")
            .withAdminPassword("ekceujoecaashtjlsgcymnrdozk")
            .withInitialCapacity(9)
            .withTags(mapOf("key6565", "cgdhmupta"))
            .withIdentity(
                new ManagedServiceIdentity()
                    .withType(ManagedServiceIdentityType.NONE)
                    .withUserAssignedIdentities(mapOf("key4522", new UserAssignedIdentity())))
            .withClusterLoginUrl("jjqhgevy")
            .withPrivateIPs(Arrays.asList("kslguxrwbwkrj"))
            .withAvailabilityZone("maseyqhlnhoiwbabcqabtedbjpip")
            .create();
    }

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
