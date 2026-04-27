
import com.azure.resourcemanager.confluent.models.Package;
import com.azure.resourcemanager.confluent.models.SCMetadataEntity;
import com.azure.resourcemanager.confluent.models.StreamGovernanceConfig;

/**
 * Samples for Environment CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-18-preview/Environment_CreateOrUpdate_MaximumSet_Gen.json
     */
    /**
     * Sample code: Environment_CreateOrUpdate_MaximumSet.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void
        environmentCreateOrUpdateMaximumSet(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.environments().define("diycvbfypirqvomdkt").withExistingOrganization("rgconfluent", "uf")
            .withKind("qhwbkvelujjbojvhrgiikildjdrqox")
            .withStreamGovernanceConfig(new StreamGovernanceConfig().withPackageProperty(Package.ESSENTIALS))
            .withMetadata(new SCMetadataEntity().withSelf("bnbnbarlsvfifpzcnsnplf")
                .withResourceName("ciadqmxlpgllibvkz").withCreatedTimestamp("ouqjivxfggaxzrsmxm")
                .withUpdatedTimestamp("ctrngbppcxdpzmp").withDeletedTimestamp("gn"))
            .create();
    }
}
