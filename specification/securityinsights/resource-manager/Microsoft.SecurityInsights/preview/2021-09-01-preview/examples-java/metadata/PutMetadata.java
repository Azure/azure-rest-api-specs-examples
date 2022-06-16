import com.azure.resourcemanager.securityinsights.models.Kind;
import com.azure.resourcemanager.securityinsights.models.MetadataAuthor;
import com.azure.resourcemanager.securityinsights.models.MetadataCategories;
import com.azure.resourcemanager.securityinsights.models.MetadataDependencies;
import com.azure.resourcemanager.securityinsights.models.MetadataSource;
import com.azure.resourcemanager.securityinsights.models.MetadataSupport;
import com.azure.resourcemanager.securityinsights.models.Operator;
import com.azure.resourcemanager.securityinsights.models.SourceKind;
import com.azure.resourcemanager.securityinsights.models.SupportTier;
import java.time.LocalDate;
import java.util.Arrays;

/** Samples for Metadata Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/metadata/PutMetadata.json
     */
    /**
     * Sample code: Create/update full metadata.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void createUpdateFullMetadata(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .metadatas()
            .define("metadataName")
            .withExistingWorkspace("myRg", "myWorkspace")
            .withContentId("c00ee137-7475-47c8-9cce-ec6f0f1bedd0")
            .withParentId(
                "/subscriptions/2e1dc338-d04d-4443-b721-037eff4fdcac/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/ruleName")
            .withVersion("1.0.0.0")
            .withKind(Kind.ANALYTICS_RULE)
            .withSource(
                new MetadataSource()
                    .withKind(SourceKind.SOLUTION)
                    .withName("Contoso Solution 1.0")
                    .withSourceId("b688a130-76f4-4a07-bf57-762222a3cadf"))
            .withAuthor(new MetadataAuthor().withName("User Name").withEmail("email@microsoft.com"))
            .withSupport(
                new MetadataSupport()
                    .withTier(SupportTier.PARTNER)
                    .withName("Microsoft")
                    .withEmail("support@microsoft.com")
                    .withLink("https://support.microsoft.com/"))
            .withDependencies(
                new MetadataDependencies()
                    .withOperator(Operator.AND)
                    .withCriteria(
                        Arrays
                            .asList(
                                new MetadataDependencies()
                                    .withOperator(Operator.OR)
                                    .withCriteria(
                                        Arrays
                                            .asList(
                                                new MetadataDependencies()
                                                    .withContentId("045d06d0-ee72-4794-aba4-cf5646e4c756")
                                                    .withKind(Kind.DATA_CONNECTOR)
                                                    .withName("Microsoft Defender for Endpoint"),
                                                new MetadataDependencies()
                                                    .withContentId("dbfcb2cc-d782-40ef-8d94-fe7af58a6f2d")
                                                    .withKind(Kind.DATA_CONNECTOR),
                                                new MetadataDependencies()
                                                    .withContentId("de4dca9b-eb37-47d6-a56f-b8b06b261593")
                                                    .withKind(Kind.DATA_CONNECTOR)
                                                    .withVersion("2.0"))),
                                new MetadataDependencies()
                                    .withContentId("31ee11cc-9989-4de8-b176-5e0ef5c4dbab")
                                    .withKind(Kind.PLAYBOOK)
                                    .withVersion("1.0"),
                                new MetadataDependencies()
                                    .withContentId("21ba424a-9438-4444-953a-7059539a7a1b")
                                    .withKind(Kind.PARSER))))
            .withCategories(
                new MetadataCategories()
                    .withDomains(Arrays.asList("Application", "Security – Insider Threat"))
                    .withVerticals(Arrays.asList("Healthcare")))
            .withProviders(Arrays.asList("Amazon", "Microsoft"))
            .withFirstPublishDate(LocalDate.parse("2021-05-18"))
            .withLastPublishDate(LocalDate.parse("2021-05-18"))
            .create();
    }
}
