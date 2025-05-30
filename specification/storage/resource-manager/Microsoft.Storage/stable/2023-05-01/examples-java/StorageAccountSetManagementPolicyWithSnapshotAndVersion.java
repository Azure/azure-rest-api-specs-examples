
import com.azure.resourcemanager.storage.fluent.models.ManagementPolicyInner;
import com.azure.resourcemanager.storage.models.DateAfterCreation;
import com.azure.resourcemanager.storage.models.DateAfterModification;
import com.azure.resourcemanager.storage.models.ManagementPolicyAction;
import com.azure.resourcemanager.storage.models.ManagementPolicyBaseBlob;
import com.azure.resourcemanager.storage.models.ManagementPolicyDefinition;
import com.azure.resourcemanager.storage.models.ManagementPolicyFilter;
import com.azure.resourcemanager.storage.models.ManagementPolicyName;
import com.azure.resourcemanager.storage.models.ManagementPolicyRule;
import com.azure.resourcemanager.storage.models.ManagementPolicySchema;
import com.azure.resourcemanager.storage.models.ManagementPolicySnapShot;
import com.azure.resourcemanager.storage.models.ManagementPolicyVersion;
import com.azure.resourcemanager.storage.models.RuleType;
import java.util.Arrays;

/**
 * Samples for ManagementPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/
     * StorageAccountSetManagementPolicyWithSnapshotAndVersion.json
     */
    /**
     * Sample code: StorageAccountSetManagementPolicyWithSnapshotAndVersion.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        storageAccountSetManagementPolicyWithSnapshotAndVersion(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getManagementPolicies()
            .createOrUpdateWithResponse("res7687", "sto9699", ManagementPolicyName.DEFAULT,
                new ManagementPolicyInner()
                    .withPolicy(
                        new ManagementPolicySchema()
                            .withRules(
                                Arrays.asList(new ManagementPolicyRule().withEnabled(true).withName("olcmtest1")
                                    .withType(RuleType.LIFECYCLE)
                                    .withDefinition(new ManagementPolicyDefinition()
                                        .withActions(new ManagementPolicyAction()
                                            .withBaseBlob(new ManagementPolicyBaseBlob()
                                                .withTierToCool(new DateAfterModification()
                                                    .withDaysAfterModificationGreaterThan(30.0F))
                                                .withTierToArchive(new DateAfterModification()
                                                    .withDaysAfterModificationGreaterThan(90.0F))
                                                .withDelete(new DateAfterModification()
                                                    .withDaysAfterModificationGreaterThan(1000.0F)))
                                            .withSnapshot(new ManagementPolicySnapShot()
                                                .withTierToCool(
                                                    new DateAfterCreation().withDaysAfterCreationGreaterThan(30f))
                                                .withTierToArchive(
                                                    new DateAfterCreation().withDaysAfterCreationGreaterThan(90f))
                                                .withDelete(
                                                    new DateAfterCreation().withDaysAfterCreationGreaterThan(1000f)))
                                            .withVersion(new ManagementPolicyVersion()
                                                .withTierToCool(
                                                    new DateAfterCreation().withDaysAfterCreationGreaterThan(30f))
                                                .withTierToArchive(
                                                    new DateAfterCreation().withDaysAfterCreationGreaterThan(90f))
                                                .withDelete(
                                                    new DateAfterCreation().withDaysAfterCreationGreaterThan(1000f))))
                                        .withFilters(new ManagementPolicyFilter()
                                            .withPrefixMatch(Arrays.asList("olcmtestcontainer1"))
                                            .withBlobTypes(Arrays.asList("blockBlob"))))))),
                com.azure.core.util.Context.NONE);
    }
}
