import com.azure.resourcemanager.storage.fluent.models.ManagementPolicyInner;
import com.azure.resourcemanager.storage.models.DateAfterModification;
import com.azure.resourcemanager.storage.models.ManagementPolicyAction;
import com.azure.resourcemanager.storage.models.ManagementPolicyBaseBlob;
import com.azure.resourcemanager.storage.models.ManagementPolicyDefinition;
import com.azure.resourcemanager.storage.models.ManagementPolicyFilter;
import com.azure.resourcemanager.storage.models.ManagementPolicyName;
import com.azure.resourcemanager.storage.models.ManagementPolicyRule;
import com.azure.resourcemanager.storage.models.ManagementPolicySchema;
import com.azure.resourcemanager.storage.models.RuleType;
import java.util.Arrays;

/** Samples for ManagementPolicies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/StorageAccountSetManagementPolicy_BaseBlobDaysAfterCreationActions.json
     */
    /**
     * Sample code: StorageAccountSetManagementPolicy_BaseBlobDaysAfterCreationActions.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountSetManagementPolicyBaseBlobDaysAfterCreationActions(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getManagementPolicies()
            .createOrUpdateWithResponse(
                "res7687",
                "sto9699",
                ManagementPolicyName.DEFAULT,
                new ManagementPolicyInner()
                    .withPolicy(
                        new ManagementPolicySchema()
                            .withRules(
                                Arrays
                                    .asList(
                                        new ManagementPolicyRule()
                                            .withEnabled(true)
                                            .withName("olcmtest1")
                                            .withType(RuleType.LIFECYCLE)
                                            .withDefinition(
                                                new ManagementPolicyDefinition()
                                                    .withActions(
                                                        new ManagementPolicyAction()
                                                            .withBaseBlob(
                                                                new ManagementPolicyBaseBlob()
                                                                    .withTierToCool(
                                                                        new DateAfterModification()
                                                                            .withDaysAfterCreationGreaterThan(30.0F))
                                                                    .withTierToArchive(
                                                                        new DateAfterModification()
                                                                            .withDaysAfterCreationGreaterThan(90.0F))
                                                                    .withDelete(
                                                                        new DateAfterModification()
                                                                            .withDaysAfterCreationGreaterThan(
                                                                                1000.0F))))
                                                    .withFilters(
                                                        new ManagementPolicyFilter()
                                                            .withPrefixMatch(Arrays.asList("olcmtestcontainer1"))
                                                            .withBlobTypes(Arrays.asList("blockBlob"))))))),
                com.azure.core.util.Context.NONE);
    }
}
