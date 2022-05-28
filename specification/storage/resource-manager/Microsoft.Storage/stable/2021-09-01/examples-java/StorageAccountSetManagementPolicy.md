Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
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
import com.azure.resourcemanager.storage.models.RuleType;
import com.azure.resourcemanager.storage.models.TagFilter;
import java.util.Arrays;

/** Samples for ManagementPolicies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/StorageAccountSetManagementPolicy.json
     */
    /**
     * Sample code: StorageAccountSetManagementPolicies.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountSetManagementPolicies(com.azure.resourcemanager.AzureResourceManager azure) {
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
                                                                            .withDaysAfterModificationGreaterThan(
                                                                                30.0f))
                                                                    .withTierToArchive(
                                                                        new DateAfterModification()
                                                                            .withDaysAfterModificationGreaterThan(
                                                                                90.0f))
                                                                    .withDelete(
                                                                        new DateAfterModification()
                                                                            .withDaysAfterModificationGreaterThan(
                                                                                1000.0f)))
                                                            .withSnapshot(
                                                                new ManagementPolicySnapShot()
                                                                    .withDelete(
                                                                        new DateAfterCreation()
                                                                            .withDaysAfterCreationGreaterThan(30f))))
                                                    .withFilters(
                                                        new ManagementPolicyFilter()
                                                            .withPrefixMatch(Arrays.asList("olcmtestcontainer1"))
                                                            .withBlobTypes(Arrays.asList("blockBlob")))),
                                        new ManagementPolicyRule()
                                            .withEnabled(true)
                                            .withName("olcmtest2")
                                            .withType(RuleType.LIFECYCLE)
                                            .withDefinition(
                                                new ManagementPolicyDefinition()
                                                    .withActions(
                                                        new ManagementPolicyAction()
                                                            .withBaseBlob(
                                                                new ManagementPolicyBaseBlob()
                                                                    .withTierToCool(
                                                                        new DateAfterModification()
                                                                            .withDaysAfterModificationGreaterThan(
                                                                                30.0f))
                                                                    .withTierToArchive(
                                                                        new DateAfterModification()
                                                                            .withDaysAfterModificationGreaterThan(
                                                                                90.0f))
                                                                    .withDelete(
                                                                        new DateAfterModification()
                                                                            .withDaysAfterModificationGreaterThan(
                                                                                1000.0f))))
                                                    .withFilters(
                                                        new ManagementPolicyFilter()
                                                            .withPrefixMatch(Arrays.asList("olcmtestcontainer2"))
                                                            .withBlobTypes(Arrays.asList("blockBlob"))
                                                            .withBlobIndexMatch(
                                                                Arrays
                                                                    .asList(
                                                                        new TagFilter()
                                                                            .withName("tag1")
                                                                            .withOp("==")
                                                                            .withValue("val1"),
                                                                        new TagFilter()
                                                                            .withName("tag2")
                                                                            .withOp("==")
                                                                            .withValue("val2")))))))),
                Context.NONE);
    }
}
```
