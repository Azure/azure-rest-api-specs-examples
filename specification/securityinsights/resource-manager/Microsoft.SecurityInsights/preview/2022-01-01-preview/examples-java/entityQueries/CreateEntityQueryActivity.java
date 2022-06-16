import com.azure.core.util.Context;
import com.azure.resourcemanager.securityinsights.models.ActivityCustomEntityQuery;
import com.azure.resourcemanager.securityinsights.models.ActivityEntityQueriesPropertiesQueryDefinitions;
import com.azure.resourcemanager.securityinsights.models.EntityType;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for EntityQueries CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/entityQueries/CreateEntityQueryActivity.json
     */
    /**
     * Sample code: Creates or updates an Activity entity query.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void createsOrUpdatesAnActivityEntityQuery(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .entityQueries()
            .createOrUpdateWithResponse(
                "myRg",
                "myWorkspace",
                "07da3cc8-c8ad-4710-a44e-334cdcb7882b",
                new ActivityCustomEntityQuery()
                    .withEtag("\"0300bf09-0000-0000-0000-5c37296e0000\"")
                    .withTitle("An account was deleted on this host")
                    .withContent("On '{{Computer}}' the account '{{TargetAccount}}' was deleted by '{{AddedBy}}'")
                    .withDescription("Account deleted on host")
                    .withQueryDefinitions(
                        new ActivityEntityQueriesPropertiesQueryDefinitions()
                            .withQuery(
                                "let GetAccountActions = (v_Host_Name:string, v_Host_NTDomain:string,"
                                    + " v_Host_DnsDomain:string, v_Host_AzureID:string, v_Host_OMSAgentID:string){\n"
                                    + "SecurityEvent\n"
                                    + "| where EventID in (4725, 4726, 4767, 4720, 4722, 4723, 4724)\n"
                                    + "// parsing for Host to handle variety of conventions coming from data\n"
                                    + "| extend Host_HostName = case(\n"
                                    + "Computer has '@', tostring(split(Computer, '@')[0]),\n"
                                    + "Computer has '\\\\', tostring(split(Computer, '\\\\')[1]),\n"
                                    + "Computer has '.', tostring(split(Computer, '.')[0]),\n"
                                    + "Computer\n"
                                    + ")\n"
                                    + "| extend Host_NTDomain = case(\n"
                                    + "Computer has '\\\\', tostring(split(Computer, '\\\\')[0]), \n"
                                    + "Computer has '.', tostring(split(Computer, '.')[-2]), \n"
                                    + "Computer\n"
                                    + ")\n"
                                    + "| extend Host_DnsDomain = case(\n"
                                    + "Computer has '\\\\', tostring(split(Computer, '\\\\')[0]), \n"
                                    + "Computer has '.', strcat_array(array_slice(split(Computer,'.'),-2,-1),'.'), \n"
                                    + "Computer\n"
                                    + ")\n"
                                    + "| where (Host_HostName =~ v_Host_Name and Host_NTDomain =~ v_Host_NTDomain) \n"
                                    + "or (Host_HostName =~ v_Host_Name and Host_DnsDomain =~ v_Host_DnsDomain) \n"
                                    + "or v_Host_AzureID =~ _ResourceId \n"
                                    + "or v_Host_OMSAgentID == SourceComputerId\n"
                                    + "| project TimeGenerated, EventID, Activity, Computer, TargetAccount,"
                                    + " TargetUserName, TargetDomainName, TargetSid, SubjectUserName, SubjectUserSid,"
                                    + " _ResourceId, SourceComputerId\n"
                                    + "| extend AddedBy = SubjectUserName\n"
                                    + "// Future support for Activities\n"
                                    + "| extend timestamp = TimeGenerated, HostCustomEntity = Computer,"
                                    + " AccountCustomEntity = TargetAccount\n"
                                    + "};\n"
                                    + "GetAccountActions('{{Host_HostName}}', '{{Host_NTDomain}}',"
                                    + " '{{Host_DnsDomain}}', '{{Host_AzureID}}', '{{Host_OMSAgentID}}')\n"
                                    + " \n"
                                    + "| where EventID == 4726 "))
                    .withInputEntityType(EntityType.HOST)
                    .withRequiredInputFieldsSets(
                        Arrays
                            .asList(
                                Arrays.asList("Host_HostName", "Host_NTDomain"),
                                Arrays.asList("Host_HostName", "Host_DnsDomain"),
                                Arrays.asList("Host_AzureID"),
                                Arrays.asList("Host_OMSAgentID")))
                    .withEntitiesFilter(mapOf("Host_OsFamily", Arrays.asList("Windows")))
                    .withEnabled(true),
                Context.NONE);
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
