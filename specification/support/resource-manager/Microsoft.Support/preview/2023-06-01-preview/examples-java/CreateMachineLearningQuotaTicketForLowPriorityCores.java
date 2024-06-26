
import com.azure.resourcemanager.support.models.Consent;
import com.azure.resourcemanager.support.models.ContactProfile;
import com.azure.resourcemanager.support.models.PreferredContactMethod;
import com.azure.resourcemanager.support.models.QuotaChangeRequest;
import com.azure.resourcemanager.support.models.QuotaTicketDetails;
import com.azure.resourcemanager.support.models.SecondaryConsent;
import com.azure.resourcemanager.support.models.SeverityLevel;
import com.azure.resourcemanager.support.models.TechnicalTicketDetails;
import com.azure.resourcemanager.support.models.UserConsent;
import java.util.Arrays;

/**
 * Samples for SupportTickets Create.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2023-06-01-preview/examples/
     * CreateMachineLearningQuotaTicketForLowPriorityCores.json
     */
    /**
     * Sample code: Create a ticket to request Quota increase for Low-priority cores for Machine Learning service.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void createATicketToRequestQuotaIncreaseForLowPriorityCoresForMachineLearningService(
        com.azure.resourcemanager.support.SupportManager manager) {
        manager.supportTickets().define("testticket").withDescription("my description").withProblemClassificationId(
            "/providers/Microsoft.Support/services/quota_service_guid/problemClassifications/machine_learning_service_problemClassification_guid")
            .withSeverity(SeverityLevel.MODERATE).withAdvancedDiagnosticConsent(Consent.YES)
            .withSupportPlanId(
                "U291cmNlOlNDTSxDbGFyaWZ5SW5zdGFsbGF0aW9uU2l0ZUlkOjcsTGluZUl0ZW1JZDo5ODY1NzIyOSxDb250cmFjdElkOjk4NjU5MTk0LFN1YnNjcmlwdGlvbklkOjc2Y2I3N2ZhLThiMTctNGVhYi05NDkzLWI2NWRhY2U5OTgxMyw=")
            .withContactDetails(new ContactProfile().withFirstName("abc").withLastName("xyz")
                .withPreferredContactMethod(PreferredContactMethod.EMAIL).withPrimaryEmailAddress("abc@contoso.com")
                .withPreferredTimeZone("Pacific Standard Time").withCountry("usa")
                .withPreferredSupportLanguage("en-US"))
            .withTitle("my title").withServiceId("/providers/Microsoft.Support/services/quota_service_guid")
            .withQuotaTicketDetails(new QuotaTicketDetails().withQuotaChangeRequestSubType("BatchAml")
                .withQuotaChangeRequestVersion("1.0").withQuotaChangeRequests(Arrays.asList(new QuotaChangeRequest()
                    .withRegion("EastUS").withPayload("{\"NewLimit\":200,\"Type\":\"LowPriority\"}"))))
            .create();
    }

    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2023-06-01-preview/examples/
     * CreateGenericQuotaTicket.json
     */
    /**
     * Sample code: Create a ticket to request Quota increase for services that do not require additional details in the
     * quotaTicketDetails object.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void
        createATicketToRequestQuotaIncreaseForServicesThatDoNotRequireAdditionalDetailsInTheQuotaTicketDetailsObject(
            com.azure.resourcemanager.support.SupportManager manager) {
        manager.supportTickets().define("testticket")
            .withDescription("Increase the maximum throughput per container limit to 10000 for account foo bar")
            .withProblemClassificationId(
                "/providers/Microsoft.Support/services/quota_service_guid/problemClassifications/cosmosdb_problemClassification_guid")
            .withSeverity(SeverityLevel.MODERATE)
            .withContactDetails(new ContactProfile().withFirstName("abc").withLastName("xyz")
                .withPreferredContactMethod(PreferredContactMethod.EMAIL).withPrimaryEmailAddress("abc@contoso.com")
                .withPreferredTimeZone("Pacific Standard Time").withCountry("usa")
                .withPreferredSupportLanguage("en-US"))
            .withTitle("my title").withServiceId("/providers/Microsoft.Support/services/quota_service_guid").create();
    }

    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2023-06-01-preview/examples/
     * CreateBatchQuotaTicketForSpecificBatchAccountForLowPriorityCores.json
     */
    /**
     * Sample code: Create a ticket to request Quota increase for Low-priority cores for a Batch account.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void createATicketToRequestQuotaIncreaseForLowPriorityCoresForABatchAccount(
        com.azure.resourcemanager.support.SupportManager manager) {
        manager.supportTickets().define("testticket").withDescription("my description").withProblemClassificationId(
            "/providers/Microsoft.Support/services/quota_service_guid/problemClassifications/batch_problemClassification_guid")
            .withSeverity(SeverityLevel.MODERATE).withAdvancedDiagnosticConsent(Consent.YES)
            .withSupportPlanId(
                "U291cmNlOlNDTSxDbGFyaWZ5SW5zdGFsbGF0aW9uU2l0ZUlkOjcsTGluZUl0ZW1JZDo5ODY1NzIyOSxDb250cmFjdElkOjk4NjU5MTk0LFN1YnNjcmlwdGlvbklkOjc2Y2I3N2ZhLThiMTctNGVhYi05NDkzLWI2NWRhY2U5OTgxMyw=")
            .withContactDetails(new ContactProfile().withFirstName("abc").withLastName("xyz")
                .withPreferredContactMethod(PreferredContactMethod.EMAIL).withPrimaryEmailAddress("abc@contoso.com")
                .withPreferredTimeZone("Pacific Standard Time").withCountry("usa")
                .withPreferredSupportLanguage("en-US"))
            .withTitle("my title").withServiceId("/providers/Microsoft.Support/services/quota_service_guid")
            .withQuotaTicketDetails(
                new QuotaTicketDetails().withQuotaChangeRequestSubType("Account").withQuotaChangeRequestVersion("1.0")
                    .withQuotaChangeRequests(Arrays.asList(new QuotaChangeRequest().withRegion("EastUS")
                        .withPayload("{\"AccountName\":\"test\",\"NewLimit\":200,\"Type\":\"LowPriority\"}"))))
            .create();
    }

    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2023-06-01-preview/examples/
     * CreateBillingSupportTicketForSubscription.json
     */
    /**
     * Sample code: Create a subscription scoped ticket for Billing related issues.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void createASubscriptionScopedTicketForBillingRelatedIssues(
        com.azure.resourcemanager.support.SupportManager manager) {
        manager.supportTickets().define("testticket").withDescription("my description").withProblemClassificationId(
            "/providers/Microsoft.Support/services/billing_service_guid/problemClassifications/billing_problemClassification_guid")
            .withSeverity(SeverityLevel.MODERATE)
            .withSupportPlanId(
                "U291cmNlOlNDTSxDbGFyaWZ5SW5zdGFsbGF0aW9uU2l0ZUlkOjcsTGluZUl0ZW1JZDo5ODY1NzIyOSxDb250cmFjdElkOjk4NjU5MTk0LFN1YnNjcmlwdGlvbklkOjc2Y2I3N2ZhLThiMTctNGVhYi05NDkzLWI2NWRhY2U5OTgxMyw=")
            .withContactDetails(new ContactProfile().withFirstName("abc").withLastName("xyz")
                .withPreferredContactMethod(PreferredContactMethod.EMAIL).withPrimaryEmailAddress("abc@contoso.com")
                .withPreferredTimeZone("Pacific Standard Time").withCountry("usa")
                .withPreferredSupportLanguage("en-US"))
            .withTitle("my title").withServiceId("/providers/Microsoft.Support/services/billing_service_guid")
            .withFileWorkspaceName("6f16735c-1530836f-e9970f1a-2e49-47b7-96cd-9746b83aa066").create();
    }

    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2023-06-01-preview/examples/
     * CreateSqlManagedInstanceQuotaTicket.json
     */
    /**
     * Sample code: Create a ticket to request Quota increase for Azure SQL managed instance.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void createATicketToRequestQuotaIncreaseForAzureSQLManagedInstance(
        com.azure.resourcemanager.support.SupportManager manager) {
        manager.supportTickets().define("testticket").withDescription("my description").withProblemClassificationId(
            "/providers/Microsoft.Support/services/quota_service_guid/problemClassifications/sql_managedinstance_problemClassification_guid")
            .withSeverity(SeverityLevel.MODERATE).withAdvancedDiagnosticConsent(Consent.YES)
            .withSupportPlanId(
                "U291cmNlOlNDTSxDbGFyaWZ5SW5zdGFsbGF0aW9uU2l0ZUlkOjcsTGluZUl0ZW1JZDo5ODY1NzIyOSxDb250cmFjdElkOjk4NjU5MTk0LFN1YnNjcmlwdGlvbklkOjc2Y2I3N2ZhLThiMTctNGVhYi05NDkzLWI2NWRhY2U5OTgxMyw=")
            .withContactDetails(new ContactProfile().withFirstName("abc").withLastName("xyz")
                .withPreferredContactMethod(PreferredContactMethod.EMAIL).withPrimaryEmailAddress("abc@contoso.com")
                .withPreferredTimeZone("Pacific Standard Time").withCountry("usa")
                .withPreferredSupportLanguage("en-US"))
            .withTitle("my title").withServiceId("/providers/Microsoft.Support/services/quota_service_guid")
            .withQuotaTicketDetails(
                new QuotaTicketDetails().withQuotaChangeRequestSubType("SQLMI").withQuotaChangeRequestVersion("1.0")
                    .withQuotaChangeRequests(Arrays.asList(
                        new QuotaChangeRequest().withRegion("EastUS")
                            .withPayload("{\"NewLimit\":200, \"Metadata\":null, \"Type\":\"vCore\"}"),
                        new QuotaChangeRequest().withRegion("EastUS")
                            .withPayload("{\"NewLimit\":200, \"Metadata\":null, \"Type\":\"Subnet\"}"))))
            .create();
    }

    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2023-06-01-preview/examples/
     * CreateBatchQuotaTicketForSpecificBatchAccountForPools.json
     */
    /**
     * Sample code: Create a ticket to request Quota increase for Pools for a Batch account.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void createATicketToRequestQuotaIncreaseForPoolsForABatchAccount(
        com.azure.resourcemanager.support.SupportManager manager) {
        manager.supportTickets().define("testticket").withDescription("my description").withProblemClassificationId(
            "/providers/Microsoft.Support/services/quota_service_guid/problemClassifications/batch_problemClassification_guid")
            .withSeverity(SeverityLevel.MODERATE).withAdvancedDiagnosticConsent(Consent.YES)
            .withSupportPlanId(
                "U291cmNlOlNDTSxDbGFyaWZ5SW5zdGFsbGF0aW9uU2l0ZUlkOjcsTGluZUl0ZW1JZDo5ODY1NzIyOSxDb250cmFjdElkOjk4NjU5MTk0LFN1YnNjcmlwdGlvbklkOjc2Y2I3N2ZhLThiMTctNGVhYi05NDkzLWI2NWRhY2U5OTgxMyw=")
            .withContactDetails(new ContactProfile().withFirstName("abc").withLastName("xyz")
                .withPreferredContactMethod(PreferredContactMethod.EMAIL).withPrimaryEmailAddress("abc@contoso.com")
                .withPreferredTimeZone("Pacific Standard Time").withCountry("usa")
                .withPreferredSupportLanguage("en-US"))
            .withTitle(
                "my title")
            .withServiceId("/providers/Microsoft.Support/services/quota_service_guid")
            .withQuotaTicketDetails(
                new QuotaTicketDetails().withQuotaChangeRequestSubType("Account").withQuotaChangeRequestVersion("1.0")
                    .withQuotaChangeRequests(Arrays.asList(new QuotaChangeRequest().withRegion("EastUS")
                        .withPayload("{\"AccountName\":\"test\",\"NewLimit\":200,\"Type\":\"Pools\"}"))))
            .create();
    }

    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2023-06-01-preview/examples/
     * CreateBatchQuotaTicketForSpecificBatchAccountForDedicatedCores.json
     */
    /**
     * Sample code: Create a ticket to request Quota increase for specific VM family cores for a Batch account.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void createATicketToRequestQuotaIncreaseForSpecificVMFamilyCoresForABatchAccount(
        com.azure.resourcemanager.support.SupportManager manager) {
        manager.supportTickets().define("testticket").withDescription("my description").withProblemClassificationId(
            "/providers/Microsoft.Support/services/quota_service_guid/problemClassifications/batch_problemClassification_guid")
            .withSeverity(SeverityLevel.MODERATE).withAdvancedDiagnosticConsent(Consent.YES)
            .withSupportPlanId(
                "U291cmNlOlNDTSxDbGFyaWZ5SW5zdGFsbGF0aW9uU2l0ZUlkOjcsTGluZUl0ZW1JZDo5ODY1NzIyOSxDb250cmFjdElkOjk4NjU5MTk0LFN1YnNjcmlwdGlvbklkOjc2Y2I3N2ZhLThiMTctNGVhYi05NDkzLWI2NWRhY2U5OTgxMyw=")
            .withContactDetails(new ContactProfile().withFirstName("abc").withLastName("xyz")
                .withPreferredContactMethod(PreferredContactMethod.EMAIL).withPrimaryEmailAddress("abc@contoso.com")
                .withPreferredTimeZone("Pacific Standard Time").withCountry("usa")
                .withPreferredSupportLanguage("en-US"))
            .withTitle("my title").withServiceId("/providers/Microsoft.Support/services/quota_service_guid")
            .withQuotaTicketDetails(new QuotaTicketDetails().withQuotaChangeRequestSubType("Account")
                .withQuotaChangeRequestVersion("1.0")
                .withQuotaChangeRequests(Arrays.asList(new QuotaChangeRequest().withRegion("EastUS").withPayload(
                    "{\"AccountName\":\"test\",\"VMFamily\":\"standardA0_A7Family\",\"NewLimit\":200,\"Type\":\"Dedicated\"}"))))
            .create();
    }

    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2023-06-01-preview/examples/
     * CreateSqlDatabaseQuotaTicketForServers.json
     */
    /**
     * Sample code: Create a ticket to request Quota increase for Servers for SQL Database.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void createATicketToRequestQuotaIncreaseForServersForSQLDatabase(
        com.azure.resourcemanager.support.SupportManager manager) {
        manager.supportTickets().define("testticket").withDescription("my description").withProblemClassificationId(
            "/providers/Microsoft.Support/services/quota_service_guid/problemClassifications/sql_database_problemClassification_guid")
            .withSeverity(SeverityLevel.MODERATE).withAdvancedDiagnosticConsent(Consent.YES)
            .withSupportPlanId(
                "U291cmNlOlNDTSxDbGFyaWZ5SW5zdGFsbGF0aW9uU2l0ZUlkOjcsTGluZUl0ZW1JZDo5ODY1NzIyOSxDb250cmFjdElkOjk4NjU5MTk0LFN1YnNjcmlwdGlvbklkOjc2Y2I3N2ZhLThiMTctNGVhYi05NDkzLWI2NWRhY2U5OTgxMyw=")
            .withContactDetails(new ContactProfile().withFirstName("abc").withLastName("xyz")
                .withPreferredContactMethod(PreferredContactMethod.EMAIL).withPrimaryEmailAddress("abc@contoso.com")
                .withPreferredTimeZone("Pacific Standard Time").withCountry("usa")
                .withPreferredSupportLanguage("en-US"))
            .withTitle("my title").withServiceId("/providers/Microsoft.Support/services/quota_service_guid")
            .withQuotaTicketDetails(new QuotaTicketDetails().withQuotaChangeRequestSubType("Servers")
                .withQuotaChangeRequestVersion("1.0").withQuotaChangeRequests(
                    Arrays.asList(new QuotaChangeRequest().withRegion("EastUS").withPayload("{\"NewLimit\":200}"))))
            .create();
    }

    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2023-06-01-preview/examples/
     * CreateTechnicalSupportTicketForSubscription.json
     */
    /**
     * Sample code: Create a subscription scoped ticket for Technical issue related to a specific resource.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void createASubscriptionScopedTicketForTechnicalIssueRelatedToASpecificResource(
        com.azure.resourcemanager.support.SupportManager manager) {
        manager.supportTickets().define("testticket").withDescription("my description").withProblemClassificationId(
            "/providers/Microsoft.Support/services/virtual_machine_running_linux_service_guid/problemClassifications/problemClassification_guid")
            .withSeverity(SeverityLevel.MODERATE).withAdvancedDiagnosticConsent(Consent.YES)
            .withProblemScopingQuestions(
                "{\"articleId\":\"076846c1-4c0b-4b21-91c6-1a30246b3867\",\"scopingDetails\":[{\"question\":\"When did the problem begin?\",\"controlId\":\"problem_start_time\",\"orderId\":1,\"inputType\":\"static\",\"answer\":{\"displayValue\":\"2023-08-31T18:55:00.739Z\",\"value\":\"2023-08-31T18:55:00.739Z\",\"type\":\"datetime\"}},{\"question\":\"API Type of the Cosmos DB account\",\"controlId\":\"api_type\",\"orderId\":2,\"inputType\":\"static\",\"answer\":{\"displayValue\":\"Table\",\"value\":\"tables\",\"type\":\"string\"}},{\"question\":\"Table name\",\"controlId\":\"collection_name_table\",\"orderId\":11,\"inputType\":\"nonstatic\",\"answer\":{\"displayValue\":\"Select Table Name\",\"value\":\"dont_know_answer\",\"type\":\"string\"}},{\"question\":\"Provide additional details about the issue you're facing\",\"controlId\":\"problem_description\",\"orderId\":12,\"inputType\":\"nonstatic\",\"answer\":{\"displayValue\":\"test ticket, please ignore and close\",\"value\":\"test ticket, please ignore and close\",\"type\":\"string\"}}]}")
            .withSupportPlanId(
                "U291cmNlOlNDTSxDbGFyaWZ5SW5zdGFsbGF0aW9uU2l0ZUlkOjcsTGluZUl0ZW1JZDo5ODY1NzIyOSxDb250cmFjdElkOjk4NjU5MTk0LFN1YnNjcmlwdGlvbklkOjc2Y2I3N2ZhLThiMTctNGVhYi05NDkzLWI2NWRhY2U5OTgxMyw=")
            .withContactDetails(new ContactProfile().withFirstName("abc").withLastName("xyz")
                .withPreferredContactMethod(PreferredContactMethod.EMAIL).withPrimaryEmailAddress("abc@contoso.com")
                .withPreferredTimeZone("Pacific Standard Time").withCountry("usa")
                .withPreferredSupportLanguage("en-US"))
            .withTitle("my title")
            .withServiceId("/providers/Microsoft.Support/services/cddd3eb5-1830-b494-44fd-782f691479dc")
            .withFileWorkspaceName("6f16735c-1530836f-e9970f1a-2e49-47b7-96cd-9746b83aa066")
            .withTechnicalTicketDetails(new TechnicalTicketDetails().withResourceId(
                "/subscriptions/subid/resourceGroups/test/providers/Microsoft.Compute/virtualMachines/testserver"))
            .withSecondaryConsent(Arrays.asList(
                new SecondaryConsent().withUserConsent(UserConsent.YES).withType("virtualmachinerunninglinuxservice")))
            .create();
    }

    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2023-06-01-preview/examples/
     * CreateMachineLearningQuotaTicketForDedicatedCores.json
     */
    /**
     * Sample code: Create a ticket to request Quota increase for specific VM family cores for Machine Learning service.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void createATicketToRequestQuotaIncreaseForSpecificVMFamilyCoresForMachineLearningService(
        com.azure.resourcemanager.support.SupportManager manager) {
        manager.supportTickets().define("testticket").withDescription("my description").withProblemClassificationId(
            "/providers/Microsoft.Support/services/quota_service_guid/problemClassifications/machine_learning_service_problemClassification_guid")
            .withSeverity(SeverityLevel.MODERATE).withAdvancedDiagnosticConsent(Consent.YES)
            .withSupportPlanId(
                "U291cmNlOlNDTSxDbGFyaWZ5SW5zdGFsbGF0aW9uU2l0ZUlkOjcsTGluZUl0ZW1JZDo5ODY1NzIyOSxDb250cmFjdElkOjk4NjU5MTk0LFN1YnNjcmlwdGlvbklkOjc2Y2I3N2ZhLThiMTctNGVhYi05NDkzLWI2NWRhY2U5OTgxMyw=")
            .withContactDetails(new ContactProfile().withFirstName("abc").withLastName("xyz")
                .withPreferredContactMethod(PreferredContactMethod.EMAIL).withPrimaryEmailAddress("abc@contoso.com")
                .withPreferredTimeZone("Pacific Standard Time").withCountry("usa")
                .withPreferredSupportLanguage("en-US"))
            .withTitle("my title").withServiceId("/providers/Microsoft.Support/services/quota_service_guid")
            .withQuotaTicketDetails(new QuotaTicketDetails().withQuotaChangeRequestSubType("BatchAml")
                .withQuotaChangeRequestVersion("1.0")
                .withQuotaChangeRequests(Arrays.asList(new QuotaChangeRequest().withRegion("EastUS")
                    .withPayload("{\"VMFamily\":\"standardA0_A7Family\",\"NewLimit\":200,\"Type\":\"Dedicated\"}"))))
            .create();
    }

    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2023-06-01-preview/examples/
     * CreateSqlDatawarehouseQuotaTicketForServers.json
     */
    /**
     * Sample code: Create a ticket to request Quota increase for Servers for Azure Synapse Analytics.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void createATicketToRequestQuotaIncreaseForServersForAzureSynapseAnalytics(
        com.azure.resourcemanager.support.SupportManager manager) {
        manager.supportTickets().define("testticket").withDescription("my description").withProblemClassificationId(
            "/providers/Microsoft.Support/services/quota_service_guid/problemClassifications/sql_datawarehouse_problemClassification_guid")
            .withSeverity(SeverityLevel.MODERATE).withAdvancedDiagnosticConsent(Consent.YES)
            .withSupportPlanId(
                "U291cmNlOlNDTSxDbGFyaWZ5SW5zdGFsbGF0aW9uU2l0ZUlkOjcsTGluZUl0ZW1JZDo5ODY1NzIyOSxDb250cmFjdElkOjk4NjU5MTk0LFN1YnNjcmlwdGlvbklkOjc2Y2I3N2ZhLThiMTctNGVhYi05NDkzLWI2NWRhY2U5OTgxMyw=")
            .withContactDetails(new ContactProfile().withFirstName("abc").withLastName("xyz")
                .withPreferredContactMethod(PreferredContactMethod.EMAIL).withPrimaryEmailAddress("abc@contoso.com")
                .withPreferredTimeZone("Pacific Standard Time").withCountry("usa")
                .withPreferredSupportLanguage("en-US"))
            .withTitle("my title").withServiceId("/providers/Microsoft.Support/services/quota_service_guid")
            .withQuotaTicketDetails(new QuotaTicketDetails().withQuotaChangeRequestSubType("Servers")
                .withQuotaChangeRequestVersion("1.0").withQuotaChangeRequests(
                    Arrays.asList(new QuotaChangeRequest().withRegion("EastUS").withPayload("{\"NewLimit\":200}"))))
            .create();
    }

    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2023-06-01-preview/examples/
     * CreateCoresQuotaTicketForSubscription.json
     */
    /**
     * Sample code: Create a ticket to request Quota increase for Compute VM Cores.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void
        createATicketToRequestQuotaIncreaseForComputeVMCores(com.azure.resourcemanager.support.SupportManager manager) {
        manager.supportTickets().define("testticket").withDescription("my description").withProblemClassificationId(
            "/providers/Microsoft.Support/services/quota_service_guid/problemClassifications/cores_problemClassification_guid")
            .withSeverity(SeverityLevel.MODERATE).withAdvancedDiagnosticConsent(Consent.YES)
            .withSupportPlanId(
                "U291cmNlOlNDTSxDbGFyaWZ5SW5zdGFsbGF0aW9uU2l0ZUlkOjcsTGluZUl0ZW1JZDo5ODY1NzIyOSxDb250cmFjdElkOjk4NjU5MTk0LFN1YnNjcmlwdGlvbklkOjc2Y2I3N2ZhLThiMTctNGVhYi05NDkzLWI2NWRhY2U5OTgxMyw=")
            .withContactDetails(new ContactProfile().withFirstName("abc").withLastName("xyz")
                .withPreferredContactMethod(PreferredContactMethod.EMAIL).withPrimaryEmailAddress("abc@contoso.com")
                .withPreferredTimeZone("Pacific Standard Time").withCountry("usa")
                .withPreferredSupportLanguage("en-US"))
            .withTitle("my title").withServiceId("/providers/Microsoft.Support/services/quota_service_guid")
            .withQuotaTicketDetails(new QuotaTicketDetails().withQuotaChangeRequestVersion("1.0")
                .withQuotaChangeRequests(Arrays.asList(new QuotaChangeRequest().withRegion("EastUS")
                    .withPayload("{\"SKU\":\"DSv3 Series\",\"NewLimit\":104}"))))
            .create();
    }
}
